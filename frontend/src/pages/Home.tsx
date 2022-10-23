import React from 'react';
import { toast } from "react-hot-toast"
import * as runtime from '@wailsapp/runtime'

export default function Home() {
    return (
        <>
            <h1>Hello World</h1>
            <button onClick={() => toast("Hello World")}>Toast</button>
            <button onClick={runtime.WindowIsMaximised}>Maximise</button>
        </>
    )
}