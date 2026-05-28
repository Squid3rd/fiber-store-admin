import { Navigate, Route, Routes } from 'react-router-dom'
import { NotFoundPage } from '../pages/NotFoundPage/NotFoundPage'
import { ProductPage } from '../pages/ProductsPage/ProductPage'
import { LoginPage } from '../pages/Login/LoginPage'

export function AppRoutes() {
  return (
    <Routes>
      <Route path="/" element={<Navigate to="/products" replace />} />
      <Route path="/products" element={<ProductPage />} />
      <Route path="/login" element={<LoginPage />} />
      <Route path="*" element={<NotFoundPage />} />
    </Routes>
  )
}