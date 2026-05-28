import { useState } from "react"
import { useProductQuery } from "../../features/products/hooks/userProductQuery"
import { ProductTable } from "../../features/products/components/ProductTable"
import { Productform } from "../../features/products/components/ProductForm"


export function ProductPage() {
    const [page] = useState(1)
    const [limit] = useState(20)


    const q = useProductQuery({ page, limit })


    return (
        <div>
            <section>
                <h2 style={{ marginTop: 12 }}>Product</h2>
                {q.isLoading ? <div>Loading...</div> : null}
                {q.isError ? <div>Error: {q.error?.message}</div> : null}
                {q.isSuccess ? <ProductTable products={q.data.data} /> : null}
                {q.data && q.data.data.length === 0 ? <div>No products found</div> : null}
            </section>

            <section>
                <h3>Product Form</h3>
                <Productform />
            </section>
        </div>
    )
}