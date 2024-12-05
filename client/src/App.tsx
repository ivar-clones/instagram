import "./App.css";
import { Button } from "@/components/ui/button";
import { ModeToggle } from "./components/mode-toggle";

function App() {
  return (
    <>
      <div>
        <Button>Click me</Button>
      </div>
      <ModeToggle />
    </>
  );
}

export default App;
