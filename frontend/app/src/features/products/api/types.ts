export type ProductStatus = "active" | "inactive";

export type Product = {
  id: string;
  name: string;
  stock: number;
  price: number;
  createdAt: string;
  updatedAt: string;
};

export type ProductCreateInput = {
  name: string;
  price: number;
  stock: number;
};

export type PaginateMeta = {
  page: number;
  limit: number;
  total: number;
};

export type ListProductsResponse = {
  data: Product[];
  meta: PaginateMeta;
};
