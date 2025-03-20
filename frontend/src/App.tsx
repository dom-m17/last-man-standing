import './App.css'
import { Button } from "@/components/ui/Button"
import Header from "@/components/ui/Header"

function App() {

  return (
    <>
      <Header></Header>
      <div className="flex flex-col items-center justify-center min-h-svh">
        <Button>Click me</Button>
      </div>
    </>
  )
}

export default App
