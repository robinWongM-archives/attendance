import Home from "./Home";
import React from "react";
import Nav from "./Nav";

function App() {
  return (
    <div className="w-screen h-screen flex bg-gray-900">
      <Nav />
      <Home />
    </div>
  );
}

export default App;
