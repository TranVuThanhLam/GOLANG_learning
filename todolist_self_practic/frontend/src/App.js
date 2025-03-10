import React, { useState, useEffect } from "react";
import axios from "axios";
import bootstrap from "bootstrap/dist/css/bootstrap.min.css";
import Todolist from "./components/Todolist";
import Layout from "./pages/main.layout";
function App() {
  return (
    <Layout>
      <div>
        <div className="App my-5"></div>
        <div className="container">
          <h1>To-Do List</h1>
          <Todolist />
        </div>
      </div>
    </Layout>
  );
}

export default App;
