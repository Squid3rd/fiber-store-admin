import type { Product } from "../api/types";

export function ProductTable({ products }: { products: Product[] }) {
    return (
        <table style={{ width: "100%" }}>
            <thead>
                <tr style={{ textAlign: "left" }}>
                    <th>Name</th>
                    <th>Price</th>
                    <th>Stock</th>
                </tr>
            </thead>
            <tbody>
                {products.map((product) => (
                    <tr key={product.id}>
                        <td>{product.name}</td>
                        <td>{product.price}</td>
                        <td>{product.stock}</td>
                    </tr>
                ))}
            </tbody>
        </table>
    );
}
