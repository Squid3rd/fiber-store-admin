import { useMutation, useQueryClient } from "@tanstack/react-query";
import type { ProductCreateInput } from "../api/types";
import { createProduct } from "../api/productsApi";

export function useCreateProduct() {
  const qc = useQueryClient();

  return useMutation({
    mutationFn: (payload: ProductCreateInput) => createProduct(payload),
    onSuccess: async () => {
      await qc.invalidateQueries({
        queryKey: ["products"],
      });
    },
  });
}
