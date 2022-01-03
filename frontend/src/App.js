import './App.css';
import React, { useState } from "react";

function App() {
  const [test, setTest] = useState([]);

  const Greet = (data) => {
    window.go.main.App.Greet(test).then(data => {
      document.getElementById("result").innerText = data;
    });
  }

  return (
    <div>
      <div className="logo"></div>
      <h1 className="text-3xl font-bold underline">
        Hello World!
      </h1>
      <div className="result" id="result">Please enter your name below 👇
          <div className="input-box" id="input" data-wails-no-drag>
          <input className="input" id="name" type="text" autoComplete="off" onChange={e => { setTest(e.target.value) }}></input>
          <button className="btn" onClick={() => Greet()}>Greet</button>
        </div>
      </div>
    </div>
  );
}

export default App;
