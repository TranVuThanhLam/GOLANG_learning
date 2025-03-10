function Todo({ todo, toggleStatus, deleteTodo }) {
  return (
    <>
      <li
        className="row list-group-item d-flex align-items-center"
        key={todo.id}
      >
        <h4
          className="col"
          style={{
            textDecoration: todo.status ? "line-through" : "none",
          }}
        >
          {todo.id}
        </h4>
        <h4
          className="col"
          style={{
            textDecoration: todo.status ? "line-through" : "none",
          }}
        >
          {todo.title}
        </h4>
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
    </>
  );
}
export default Todo;
