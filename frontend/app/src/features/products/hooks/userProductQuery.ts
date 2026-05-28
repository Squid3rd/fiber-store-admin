import { useQuery } from "@tanstack/react-query";
import { listProducts } from "../api/productsApi";

export function useProductQuery(params?: { page: number; limit: number }) {
  return useQuery({
    queryKey: ["products", params.page, params.limit],
    queryFn: () => listProducts(params),
  });
}
