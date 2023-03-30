package messaging

// EmailTemplateBody stores email body templates.
var EmailTemplateBody = map[string]string{
	"en/registration_confirmation": `<html>
  <body>
    <p>
      Please confirm your registration by clicking this
      <a href="{{ .registration_url }}/ack/{{ .registration_id }}">link</a>
      and providing the registration code <b><code>{{ .registration_code }}</code></b>
      within the next 45 minutes. If you haven't done so, please re-register.
    </p>

    <p>The registation metadata follows:</p>
    <ul style="list-style-type: disc">
      <li>Session ID: {{ .session_id }}</li>
      <li>Request ID: {{ .request_id }}</li>
      <li>Username: <code>{{ .username }}</code></li>
      <li>Email: <code>{{ .email }}</code></li>
      <li>IP Address: <code>{{ .src_ip }}</code></li>
      <li>Timestamp: {{ .timestamp }}</li>
    </ul>
  </body>
</html>`,
	"en/registration_ready": `<html>
  <body>
    <p>
      The following user successfully registered with the portal.
      Please use management interface to approve or decline the registration.
    </p>

    <p>The registation metadata follows:</p>
    <ul style="list-style-type: disc">
      <li>Registration ID: {{ .registration_id }}</li>
      <li>Registration URL: <code>{{ .registration_url }}</code></li>
      <li>Session ID: {{ .session_id }}</li>
      <li>Request ID: {{ .request_id }}</li>
      <li>Username: <code>{{ .username }}</code></li>
      <li>Email: <code>{{ .email }}</code></li>
      <li>IP Address: <code>{{ .src_ip }}</code></li>
      <li>Timestamp: {{ .timestamp }}</li>
    </ul>
  </body>
</html>`,
	"en/registration_verdict": `<html>
  <body>
    <p>
    {{- if eq .verdict "approved" -}}
      Your registration has been approved.
      You may now login with the username or email
      address below.
    {{- else -}}
      Your registration has been declined.
    {{- end -}}
    </p>
    <p>The registation metadata follows:</p>
    <ul style="list-style-type: disc">
      <li>Username: <code>{{ .username }}</code></li>
      <li>Email: <code>{{ .email }}</code></li>
      <li>Timestamp: {{ .timestamp }}</li>
    </ul>
  </body>
</html>`,
}
