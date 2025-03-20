import './App.css'
import { Button } from "@/components/ui/Buttontmp"
import Header from "@/components/ui/Header_tmp"

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
