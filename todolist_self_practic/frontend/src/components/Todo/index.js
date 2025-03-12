function Todo({ todo, toggleStatus, confirmDelete }) {
  return (
    <>
      <tr key={todo.id}>
        <th scope="row">{todo.id}</th>
        <td>{todo.title}</td>
        <td></td> {/* Cột spacer trống */}
        <td>
          <button
            className="btn btn-outline-primary mr-2"
            onClick={() => toggleStatus(todo.id, todo.status)}
          >
            {todo.status ? "Undo" : "Done"}
          </button>
          <button
            data-bs-toggle="modal"
            data-bs-target="#staticBackdrop"
            className="btn btn-outline-danger mx-2"
            onClick={() => confirmDelete(todo.id)}
          >
            Delete
          </button>
        </td>
      </tr>
    </>
  );
}
export default Todo;
