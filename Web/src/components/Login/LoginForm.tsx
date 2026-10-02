interface LoginFormProps {
  username: string
  password: string
  loginError: string
  onUsernameChange: (username: string) => void
  onPasswordChange: (password: string) => void
  onLogin: () => void
}

function LoginForm({
  username,
  password,
  loginError,
  onUsernameChange,
  onPasswordChange,
  onLogin,
}: LoginFormProps) {
  return (
    <div className="login-page">
      <div className="login-card">
        <h1>NOX</h1>
        <p>Sign in to continue</p>

        <input
          type="text"
          placeholder="Username"
          value={username}
          onChange={(event) =>
            onUsernameChange(event.target.value)
          }
        />

        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(event) =>
            onPasswordChange(event.target.value)
          }
          onKeyDown={(event) => {
            if (event.key === 'Enter') {
              onLogin()
            }
          }}
        />

        {loginError !== '' && (
          <p className="login-error">{loginError}</p>
        )}

        <button onClick={onLogin}>
          Login
        </button>
      </div>
    </div>
  )
}

export default LoginForm