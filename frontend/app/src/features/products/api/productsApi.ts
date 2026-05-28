import { http } from "../../../shared/api/http";
import type { ListProductsResponse, Product, ProductCreateInput } from "./types";

export function listProducts(params?: { page: number; limit: number }) {
  const { page = 1, limit = 10 } = params ?? {};

  return http<ListProductsResponse>(`/products?page=${page}&limit=${limit}`);
}

export  function createProduct(payload: ProductCreateInput){
    return http<Product>(
        `/products`,
        {
            method: 'POST',
            body: JSON.stringify(payload)
        }
    )
}
