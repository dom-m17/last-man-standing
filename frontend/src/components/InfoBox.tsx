import { ReactNode } from 'react'

type InfoBoxProps = {
    mode : 'hint' | 'warning'
    children: ReactNode
}

export default function InfoBox({mode, children}: InfoBoxProps) {
    return <aside>
        {mode === 'warning' ? <h2>Warning</h2> : null}
        <p>{children}</p>
    </aside>
}