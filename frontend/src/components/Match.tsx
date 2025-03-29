// import { useState } from 'react';

interface matchProps {
    children?: string
    homeTeam: string
    awayTeam: string
    onSelect: (team: string) => void
    selectedTeam: string
  }

export default function Match({ homeTeam, awayTeam, onSelect, selectedTeam, ...props }: matchProps) {
  
return (
    <span className="match" {...props}>
    <button className={selectedTeam === homeTeam ? "active" : ""} onClick={() => {onSelect(homeTeam)}}>{homeTeam}</button>
    <div>v</div>
    <button className={selectedTeam === awayTeam ? "active" : ""} onClick={() => {onSelect(awayTeam)}}>{awayTeam}</button>
    </span>
)
}