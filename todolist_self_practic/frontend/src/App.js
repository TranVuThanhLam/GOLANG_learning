import React, { useState, useEffect } from "react";
import axios from "axios";
import bootstrap from "bootstrap/dist/css/bootstrap.min.css";
import Todolist from "./components/Todolist";
import Layout from "./pages/main.layout";
function App() {
  return (
    <Layout>
      <div>
        <Todolist />
      </div>
    </Layout>
  );
}

export default App;
