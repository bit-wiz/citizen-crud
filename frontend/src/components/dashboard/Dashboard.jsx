import {
  Tabs,
  TabsContent,
  TabsList,
  TabsTrigger,
} from "@/src/components/ui/tabs"
import AllData from "./AllData"

export const metadata = {
  title: "Dashboard",
  description: "Example dashboard app built using the components.",
}

export default function DashboardPage() {

  const tabs = [
    {
      value: "alldata",
      component : <AllData/>
    }
  ]

  return (
    <>
      <div className="hidden flex-col md:flex">
        <div className="border-b">
        </div>
        <div className="flex-1 space-y-4 p-8 pt-6">
          <div className="flex items-center justify-between space-y-2">
            <h2 className="text-3xl font-bold tracking-tight">Dashboard</h2>
          </div>
          <Tabs defaultValue="alldata" className="space-y-4">
            <TabsList>

              { tabs.map((tab, i) => (
                <div key={i} >
                <TabsTrigger value={tab.value} >
                  {tab.value.toUpperCase()}
                </TabsTrigger>
                </div>
              ))}

            </TabsList>
            { tabs.map((tab, i) => (

              <div key={i}>
                <TabsContent value={tab.value} className="space-y-4">
                  {tab.component}
                </TabsContent>
              </div>
            ))

            }
          </Tabs>
        </div>
      </div>
    </>
  )
}