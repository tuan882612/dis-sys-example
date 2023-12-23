package smtp

const (
	AuthTemplate = `<!DOCTYPE html>
	<html>
	<head>
		<title>Auth Token</title>
	</head>
	<body>
		<div class="email-container">
			<div class="email-content">
				<p>
					This token is valid for <strong>3 minutes</strong>. 
					Please use it to complete your authentication process.
				</p>
			</div>
			<div class="footer">
				<p>
					If you did not request this token, please ignore this email or contact support for assistance.
				</p>
			</div>
		</div>
	</body>
	</html>`
)
