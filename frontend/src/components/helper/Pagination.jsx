import {
    Pagination,
    PaginationContent,
    PaginationEllipsis,
    PaginationItem,
    PaginationLink,
    PaginationNext,
    PaginationPrevious,
  } from "@/src/components/ui/pagination"
import { Button } from "../ui/button"

  export function PaginationData({page, setPage}) {
    return (
      <Pagination>
        <PaginationContent>
          <PaginationItem>
            <Button  variant="outline" onClick={() => setPage(page > 1 ? page - 1 : 1)}>
                <PaginationPrevious />
            </Button>
          </PaginationItem>
          <PaginationItem>
            <Button variant="outline" onClick={() => setPage(1)}>
                <PaginationLink>{page}</PaginationLink>
            </Button>
          </PaginationItem>
          <PaginationItem>
            <PaginationEllipsis />
          </PaginationItem>
          <PaginationItem>
            <Button variant="outline" onClick={() => setPage(page + 1)} >
                <PaginationNext />
            </Button>
          </PaginationItem>
        </PaginationContent>
      </Pagination>
    )
  }
