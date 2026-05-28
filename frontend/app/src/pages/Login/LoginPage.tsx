
export function LoginPage() {
    return (
        <div style={{display: 'grid', gap: 12}}>
            <h2>Login</h2>
            <form style={{display: 'grid', gap: 12}}>
                <div style={{display: 'grid', gap: 4}}>
                    <label style={{textAlign: 'left'}}>Email</label>
                    <input type="email" />
                </div>
                <div style={{display: 'grid', gap: 4}}>
                    <label style={{textAlign: 'left'}}>Password</label>
                    <input type="password" />
                </div>
                <button type="submit" style={{marginTop: 10, height: 40}}>Login</button>
            </form>
        </div>
    )
}