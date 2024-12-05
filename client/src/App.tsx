import "./App.css";
import { ModeToggle } from "@/components/mode-toggle";
import { SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar";
import { AppSidebar } from "@/components/app-sidebar";
import { Button } from "@/components/ui/button";

function App() {
  return (
    <>
      <SidebarProvider>
        <AppSidebar />
        <main className="w-full pt-2 pr-2 pb-2">
          <div className="flex flex-row items-center justify-between">
            <SidebarTrigger />
            <ModeToggle />
          </div>
        </main>
      </SidebarProvider>
    </>
  );
}

export default App;
