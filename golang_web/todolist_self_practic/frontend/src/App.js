import React, { useState, useEffect } from "react";
import axios from "axios";

function App() {
  const [todos, setTodos] = useState([]);
  const [title, setTitle] = useState("");

  useEffect(() => {
    axios.get("http://localhost:8080/todos").then((res) => setTodos(res.data));
  }, []);

  const addTodo = () => {
    axios.post("http://localhost:8080/todos", { title, status: false }).then((res) => {
      setTodos([...todos, res.data]);
      setTitle("");
    });
  };

  const toggleStatus = (id, status) => {
    axios.put(`http://localhost:8080/todos/${id}`, { status: !status }).then(() => {
      setTodos(todos.map(todo => (todo.ID === id ? { ...todo, status: !status } : todo)));
    });
  };

  const deleteTodo = (id) => {
    axios.delete(`http://localhost:8080/todos/${id}`).then(() => {
      setTodos(todos.filter(todo => todo.ID !== id));
    });
  };

  return (
    <div>
      <h1>To-Do List</h1>
      <input value={title} onChange={(e) => setTitle(e.target.value)} />
      <button onClick={addTodo}>Add</button>
      <ul>
        {todos.map((todo) => (
          <li key={todo.ID}>
            <span style={{ textDecoration: todo.status ? "line-through" : "none" }}>
              {todo.title}
            </span>
            <button onClick={() => toggleStatus(todo.ID, todo.status)}>
              {todo.status ? "Undo" : "Done"}
            </button>
            <button onClick={() => deleteTodo(todo.ID)}>Delete</button>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default App;
