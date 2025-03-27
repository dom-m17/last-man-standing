import { useState } from 'react'; 

interface matchProps {
    children?: string
    homeTeam: string
    awayTeam: string
  }

export default function Match({ homeTeam, awayTeam, ...props }: matchProps) {

  const [ active, setActive ] = useState(0)

  function handleClick(val: number) {
    setActive(val)
  }
  
return (
    <span className="match" {...props}>
    <button className={active==1 ? "active" : ""} onClick={() => handleClick(active == 1 ? 0 : 1)}>{homeTeam}</button>
    <div>v</div>
    <button className={active==2 ? "active" : ""} onClick={() => handleClick(active == 2 ? 0 : 2)}>{awayTeam}</button>
    </span>
)
}