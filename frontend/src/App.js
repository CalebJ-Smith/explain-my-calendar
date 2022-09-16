// App.js
import React, { useEffect, useState } from "react";
import Header from "./components/Header/Header"
import { ping } from "./api/goService";

const Small = ({ pingResponse }) => {
  return (
    <h2>ping returns: {pingResponse}</h2>
  )
}

function App() {
  const [pingResult, setPingResult] = useState('')
  useEffect(() => {
    ping().then(r => {
      console.log(r);
      setPingResult(r);
    })
  }, [])
  //<div><p>{pingResult}</p></div>
  return (
    <div className="App">
      <Header />
      <Small pingResponse={pingResult} />
    </div>
  );
}

export default App;