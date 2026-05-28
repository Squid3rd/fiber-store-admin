import { useState } from "react";
import type { ProductCreateInput } from "../api/types";
import { useCreateProduct } from "../hooks/userCreateProduct";

export function Productform() {
    const create = useCreateProduct();

    const [form, setForm] = useState<ProductCreateInput>({
        name: "",
        price: 0,
        stock: 0,
    });

    function set<K extends keyof ProductCreateInput>(
        key: K,
        value: ProductCreateInput[K],
    ) {
        setForm((prev) => ({
            ...prev,
            [key]: value,
        }));
    }

    async function handleSubmit(e: React.FormEvent) {
        e.preventDefault();
        await create.mutateAsync(form);
        setForm({
            name: "",
            price: 0,
            stock: 0,
        });
    }

    return (
        <form onSubmit={handleSubmit} style={{ display: "grid", gap: 12 }}>
            <div style={{ display: "grid", gap: 4 }}>
                <label>name</label>
                <input
                    type="text"
                    value={form.name}
                    onChange={(e) => set("name", e.target.value)}
                    required
                />
            </div>

            <div style={{ display: "grid", gap: 4 }}>
                <label htmlFor="price">Price</label>
                <input
                    type="number"
                    value={form.price}
                    onChange={(e) => set("price", Number(e.target.value))}
                    required
                />
            </div>

            <div style={{ display: "grid", gap: 4 }}>
                <label htmlFor="stock">Stock</label>
                <input
                    type="number"
                    value={form.stock}
                    onChange={(e) => set("stock", Number(e.target.value))}
                    required
                />
            </div>

            <button type="submit" disabled={create.isPending}>
                {create.isPending ? "Creating..." : "Create product"}
            </button>

            {create.isError ? (
                <div style={{ color: "crimson" }}>
                    {create.error?.message ?? "Create failed"}
                </div>
            ) : null}
        </form>
    );
}
