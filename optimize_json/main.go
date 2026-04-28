package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Condition đại diện cho điều kiện đơn (leaf node)
type Condition struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Type  string `json:"type"` // "string", "number", "boolean", ...
}

// QueryNode đại diện cho node trong cây logic
type QueryNode struct {
	Operator  string      `json:"operator,omitempty"`  // "AND"/"OR" nếu node nhóm
	Children  []QueryNode `json:"children,omitempty"`  // các node con
	Condition *Condition  `json:"condition,omitempty"` // node lá
}

// UnmarshalJSON hỗ trợ leaf node trực tiếp hoặc có key "condition" và AND/OR node
func (q *QueryNode) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	// Node nhóm
	for k, v := range raw {
		if k == "AND" || k == "OR" {
			q.Operator = k
			var children []QueryNode
			if err := json.Unmarshal(v, &children); err != nil {
				return err
			}
			q.Children = children
			return nil
		}
	}

	// Node leaf có key "condition"
	if condData, ok := raw["condition"]; ok {
		var cond Condition
		if err := json.Unmarshal(condData, &cond); err != nil {
			return err
		}
		q.Condition = &cond
		return nil
	}

	// Node leaf trực tiếp
	var cond Condition
	if err := json.Unmarshal(data, &cond); err != nil {
		return err
	}
	q.Condition = &cond
	return nil
}

// Human readable print
func prettyQuery(node QueryNode, indent int) string {
	pad := ""
	for i := 0; i < indent; i++ {
		pad += "  "
	}

	if node.Condition != nil {
		return fmt.Sprintf("%s- %s: %s (%s)\n", pad, node.Condition.Field, node.Condition.Value, node.Condition.Type)
	}

	res := fmt.Sprintf("%s%s\n", pad, node.Operator)
	for _, child := range node.Children {
		res += prettyQuery(child, indent+1)
	}
	return res
}

// Handler nhận POST request, parse JSON, in ra console
func queryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var root QueryNode
	if err := json.NewDecoder(r.Body).Decode(&root); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// In cây logic lên console
	fmt.Println("Received query JSON:")
	fmt.Println(prettyQuery(root, 0))

	// Có thể trả simple response để client biết OK
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Query received and printed on server console.\n"))
}

func main() {
	http.HandleFunc("/query", queryHandler)
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
