import { useState } from 'react';
import Match from "@/components/Match";

const matches = [
    { id: 1, homeTeam: "Arsenal", awayTeam: "Tottenham" },
    { id: 2, homeTeam: "Manchester United", awayTeam: "Chelsea" },
    { id: 3, homeTeam: "Liverpool", awayTeam: "Manchester City" },
    { id: 4, homeTeam: "Newcastle", awayTeam: "Aston Villa" },
    { id: 5, homeTeam: "West Ham", awayTeam: "Brighton" }
    ];

export default function Matches({...props}) {

    const [ selectedTeam, setSelectedTeam ] = useState("")

    return (
        <>
            {matches.map(
                      (match) => 
                      <Match
                      {...props}
                      key={match.id}
                      onSelect={setSelectedTeam}
                      selectedTeam={selectedTeam}
                      {...match}>
                      </Match>
                    )}
        </>
    )
}