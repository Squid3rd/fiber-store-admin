import { AppRoutes } from "./routes/routes";

export default function App() {
  return (
    <div style={{ maxWidth: '1200px', margin: '0 auto' }}>
      <header>
        <h1>Fiber Store Admin</h1>
        <span>With Go Api</span>
      </header>

      <main style={{ marginTop: '2rem' }}>
        <AppRoutes />
      </main>
    </div>
  )
}