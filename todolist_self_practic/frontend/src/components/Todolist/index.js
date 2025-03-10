import React, { useState, useEffect } from "react";
import axios from "axios";
import Todo from "../Todo";

function Todolist() {
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
    <>
      <div className="input-group mb-3">
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

      <div className="justify-content-between">
        <ul className="list-group">
          {todos.map((todo) => (
            <Todo
              key={todo.id}
              todo={todo}
              toggleStatus={toggleStatus}
              deleteTodo={deleteTodo}
            />
          ))}
        </ul>
      </div>
    </>
  );
}
export default Todolist;
