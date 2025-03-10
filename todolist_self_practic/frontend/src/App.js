import React, { useState, useEffect } from "react";
import axios from "axios";
import bootstrap from "bootstrap/dist/css/bootstrap.min.css";

function App() {
  const [todos, setTodos] = useState([]);
  const [title, setTitle] = useState("");

  useEffect(() => {
    axios.get("http://localhost:8080/todos").then((res) => setTodos(res.data));
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

  const deleteTodo = (id) => {
    axios.delete(`http://localhost:8080/todos/${id}`).then(() => {
      setTodos(todos.filter((todo) => todo.id !== id));
    });
  };

  return (
    <div>
      <div className="App my-5"></div>
      <div className="container">
        <h1>To-Do List</h1>
        <div className="input-group mb-3">
          <input
            className="form-control"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
          />
          <button className="btn btn-outline-primary" onClick={addTodo}>
            Add
          </button>
        </div>

        <div className="justify-content-between">
          <ul className="list-group">
            {todos.map((todo) => (
              <li
                className="row list-group-item d-flex align-items-center"
                key={todo.id}
              >
                <span
                  className="col"
                  style={{
                    textDecoration: todo.status ? "line-through" : "none",
                  }}
                >
                  {todo.id}
                </span>
                <span
                  className="col"
                  style={{
                    textDecoration: todo.status ? "line-through" : "none",
                  }}
                >
                  {todo.title}
                </span>
                <button
                  className="col btn btn-outline-primary"
                  onClick={() => toggleStatus(todo.id, todo.status)}
                >
                  {todo.status ? "Undo" : "Done"}
                </button>
                <button
                  className="col btn btn-outline-danger"
                  onClick={() => deleteTodo(todo.id)}
                >
                  Delete
                </button>
              </li>
            ))}
          </ul>
        </div>
      </div>
    </div>
  );
}

export default App;
