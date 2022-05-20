import {useState} from 'react';
import logo from './assets/images/logo-universal.png';
import {Greet} from "../wailsjs/go/main/App";

function App() {
    const [resultText, setResultText] = useState("Please enter your name below 👇");
    const [name, setName] = useState('');
    const updateName = (e: any) => setName(e.target.value);
    const updateResultText = (result: string) => setResultText(result);

    function greet() {
        Greet(name).then(updateResultText);
    }

    return (
      <div className="relative bg-white overflow-hidden">
        <h1>Hello World!</h1>
      </div>
    )
}

export default App
