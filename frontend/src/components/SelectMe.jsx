import React from 'react'
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from "@/src/components/ui/select"

const SelectMe = ({ setVal, km, arr, pay }) => {
  return (
    <>
      <Select value={pay[km]} defaultValue={arr[0]}
      onValueChange={(v) => setVal({...pay, [`${km}`]: v })}
      >
          <SelectTrigger className="w-full px-0 h-[16px] text-[13px] border-none" style={{ backgroundColor: '#fff', color: '#333' }}>
              <SelectValue />
          </SelectTrigger>
          <SelectContent>
              {arr.map((v, i) => (
                  <SelectItem
                      key={i}
                      value={v}
                  >
                      {v}
                  </SelectItem>
              ))}
          </SelectContent>
      </Select>
    </>
  )
}

export default SelectMe