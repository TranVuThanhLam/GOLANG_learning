import React, { useState, useEffect } from "react";
import axios from "axios";
import Todo from "../Todo";
import PopupDelete from "../common/PopupDelete";

function Todolist() {
  const [todos, setTodos] = useState([]);
  const [title, setTitle] = useState("");
  const [deleteId, setDeleteId] = useState(null);

  useEffect(() => {
    axios.get("http://localhost:8080/todos").then((res) => setTodos(res.data));
    if (todos.length === 0) {
      console.log("todo is empty");
    } else {
      console.log(todos);
    }
  }, []);

  const addTodo = () => {
    axios
      .post("http://localhost:8080/todos", { title, status: false })
      .then((res) => {
        setTodos([...todos, res.data]);
        setTitle("");
      });
  };

  const toggleStatus = (id, status) => {
    axios
      .put(`http://localhost:8080/todos/${id}`, { status: !status })
      .then(() => {
        setTodos(
          todos.map((todo) =>
            todo.id === id ? { ...todo, status: !status } : todo
          )
        );
      });
  };

  const confirmDelete = (id) => {
    setDeleteId(id);
  };

  const handleDelete = () => {
    if (deleteId !== null) {
      axios.delete(`http://localhost:8080/todos/${deleteId}`).then(() => {
        setTodos(todos.filter((todo) => todo.id !== deleteId));
        setDeleteId(null);
      });
    }
  };

  return (
    <>
      <div
        class="modal fade"
        id="staticBackdrop"
        data-bs-backdrop="static"
        data-bs-keyboard="false"
        tabindex="-1"
        aria-labelledby="staticBackdropLabel"
        aria-hidden="true"
      >
        <PopupDelete handleDelete={handleDelete} />
      </div>

      <div className="input-group mb-3">
        <span className="input-group-text">Add Task</span>
        <input
          className="form-control"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          onKeyDown={(e) => {
            if (e.key === "Enter") {
              addTodo();
            }
          }}
        />
        <button className="btn btn-outline-primary" onClick={addTodo}>
          Add
        </button>
      </div>

      <table className="table" style={{ width: "100%", tableLayout: "fixed" }}>
        <colgroup>
          {/* Cột đầu tiên: ID */}
          <col style={{ width: "10%" }} />
          {/* Cột thứ hai: Task */}
          <col style={{ width: "40%" }} />
          {/* Cột spacer: tự động chiếm không gian còn lại */}
          <col style={{ width: "auto" }} />
          {/* Cột cuối: Edit/Delete */}
          <col style={{ width: "20%" }} />
        </colgroup>
        <thead>
          <tr className="table-success">
            <th scope="col">#</th>
            <th scope="col">Task</th>
            <th scope="col"></th> {/* Cột spacer không có tiêu đề */}
            <th scope="col">Edit/Delete</th>
          </tr>
        </thead>
        <tbody>
          {todos.length === 0 ? (
            <tr>
              <td colSpan="4" className="text-center">
                No tasks available
              </td>
            </tr>
          ) : (
            todos.map((todo) => (
              <Todo
                key={todo.id}
                todo={todo}
                toggleStatus={toggleStatus}
                confirmDelete={confirmDelete}
              />
            ))
          )}
        </tbody>
      </table>
      <script
        src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/js/bootstrap.bundle.min.js"
        integrity="sha384-YvpcrYf0tY3lHB60NNkmXc5s9fDVZLESaAA55NDzOxhy9GkcIdslK1eN7N6jIeHz"
        crossorigin="anonymous"
      ></script>
    </>
  );
}
export default Todolist;
