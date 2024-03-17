import { Button } from "@/src/components/ui/button"
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/src/components/ui/dialog"
import { Input } from "@/src/components/ui/input"
import { conf } from "@/src/lib/constant";
import { fetchme } from "@/src/lib/fetcher";
import { useEffect } from "react";
import { useState } from "react";

export function CitizenDialog({data, type}) {
    const [payload, setPayload] = useState({})

    async function doit() {
        const url = type === 'Edit' ? `${conf['citizen']}/${data.id}` : conf['citizen']
        const method = type === 'Edit' ? "PATCH" : "POST"
        const res = await fetchme(url, method, JSON.stringify(payload))
        if(res && !res.error) {
            await Promise.resolve(r => setTimeout(r, 1000))
            window.location.reload()
        }
    }

    useEffect(() => {
        if(type === 'Edit') {
            Object.keys(data).filter(e => e!== 'id' && !e.includes('_at')).map((d, _) => {
                setPayload((prev) => ({...prev, [d]: data[d]}))
            })
        } else {
            console.log(type, data);
            data.filter(e => e!== 'id' && !e.includes('_at')).map((d, _) => {
                setPayload((prev) => ({...prev, [d]: ''}))
                // console.log();
            })
        }
    }, [])

    function Iter(varr) {
        if (Array.isArray(varr)) {
            return varr.filter(e => e !== 'id' && !e.includes('_at'));
        } else {
            return Object.keys(varr).filter(e => e !== 'id' && !e.includes('_at'));
        }
    }

    return (
        <Dialog>
            <DialogTrigger asChild>
            <Button variant="outline">{type} Profile</Button>
            </DialogTrigger>
            <DialogContent className="sm:max-w-[425px] dark">
            <DialogHeader>
                <DialogTitle>{type} profile</DialogTitle>
                <DialogDescription>
                Make changes to your profile here. Click save when you're done.
                </DialogDescription>
            </DialogHeader>
            <div className="h-[300px] overflow-y-auto py-4">
                <div className="items-center gap-4">
                {Iter(data).map((d, i) => (
                    <div key={i}>
                    <label htmlFor="name" className="text-right">
                        {d}
                    </label>
                    <Input defaultValue={type === 'Edit' ? data[d] : ''}
                    placeholder={d}
                    className="col-span-3"
                    onChange={(e) => setPayload({...payload, [d]: e.target.value})}
                    />
                    </div>
                ))}
                </div>
            </div>
            <DialogFooter>
                <Button type="submit" onClick={doit} >Save changes</Button>
            </DialogFooter>
            </DialogContent>
        </Dialog>
        );

}
