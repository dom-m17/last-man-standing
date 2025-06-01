import { ReactNode } from "react"

type imgProps = {
    img: {
        src: string
        alt: string
    }
    children: ReactNode
}

export default function Header({img, children}: imgProps) {
    return <h1>
        <img {...img} />
        {children}
    </h1>
}