import React from 'react'
import {
    Card,
    CardContent,
    CardDescription,
    CardHeader,
    CardTitle,
  } from "@/src/components/ui/card"

import { useState, useEffect } from 'react'
import { fetchit, fetchme } from '@/src/lib/fetcher'
import { Input } from '../ui/input'
import { Button } from '../ui/button'
import SelectMe from '../SelectMe'
import { conf } from '@/src/lib/constant'
import { PaginationData } from '../helper/Pagination'
import { CitizenDialog } from '../helper/CitizenDialog'

const AllData = () => {
    const [isLoading, setIsLoading] = useState(true)
    const [data, setData] = useState([])
    const [fields, setFields] = useState([])
    const [payload, setPayload] = useState({search: '', key: '', val: 'ascending'})
    const [page, setPage] = useState(1)

    async function getNew() {
        console.log('called ', payload);
        const res = await fetchme(
            `${conf['alldata']}?page=${page}&s=${payload.search}&sort=${payload.val[0]}-${payload.key}`
        )
        setData(res.data)
    }

    useEffect(() => {
        if(!isLoading) {
            getNew()
        }
    }, [page])

    useEffect(() => {
        Promise.all([
            fetchme(conf['alldata']).then((res) => {
                setData(res.data)
            }),
            fetchme(conf['fields']).then((res) => {
                setFields(res.data)
                setPayload({...payload, key: res.data[0]})
            })
        ]).then(() => {
            setIsLoading(false)
        })
    }, [])

    if(isLoading) {
        return <>Loading...</>
    }
  return (
    <>
        <Card>
            <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-2xl font-bold">
                All Data [ TOTAL records {data && data.length}]
            </CardTitle>
            </CardHeader>
                <div className='container px-5 py-4' >

                            <label className='mt-4' htmlFor="search">Search</label>
                            <Input className='mt-4' placeholder='search' type='text' onChange={(e) => setPayload({...payload, search: e.target.value})} />
                            <label className='mt-4' htmlFor="key">Key</label>
                            <SelectMe km={'key'} pay={payload} arr={fields} setVal={setPayload} />
                            <label className='mt-4' htmlFor="val">Value</label>
                            <SelectMe km={'val'} pay={payload} arr={['ascending', 'descending']} setVal={setPayload} />
                    <Button className='mt-4' onClick={getNew} >Search</Button>
                </div>
                <CardContent>
                    <CitizenDialog data={fields} type={'Add'} />
                    {data && data.map((d, i) => <SingleData key={i} data={d} />)}
                </CardContent>
                <PaginationData page={page} setPage={setPage} />
        </Card>
    </>
  )
}

function SingleData({data}) {
    async function onDelete() {
        console.log(data.id);
        const ans = confirm(`Are you sure, you want to delete user ${data.first_name || 'unknown'} ?`)
        if(ans) {
            const res = await fetchme(`${conf['citizen']}/${data.id}`, "DELETE")
            if(res && !res.error) {
                alert('deleted')
                await Promise.resolve(r => setTimeout(r, 1000))
            }
        } else {
            alert('cancelled')
        }
    }

    async function edit() {}
    return (
        <>
        <Card>
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <CardTitle className="text-xl font-medium">
            {data.first_name} {data.last_name}
            {/* a delete button */}
        </CardTitle>
        </CardHeader>
        <CardContent>
            <div className='text-xs text-muted-foreground flex justify-between items-center'>
                <div>
                    <span className="text-sm font-bold">{data.city}, {data.state}</span>
                    <p className="text-xs text-muted-foreground">
                        {data.address}
                    </p>
                </div>
                <div className="text-xs text-muted-foreground">
                    <CitizenDialog data={data} type={'Edit'} />
                </div>
                <div className="text-xs text-muted-foreground flex items-baseline">
                    <Button onClick={onDelete}>delete</Button>
                </div>
                <div className="text-xs text-muted-foreground flex items-baseline">
                    <span className="mr-2">Gender: {data.gender}</span>
                </div>
            </div>
        </CardContent>
        </Card>
        </>
    )
}

export default AllData