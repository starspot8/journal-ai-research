JOURNAL AI RESEARCH

AI-08 DEPLOYMENT SHARED RULES

---

1. PRODUCTION ARCHITECTURE

Target:

                    INTERNET
                       |
                     HTTPS
                       |
                Reverse Proxy
                       |
             +---------+---------+
             |                   |
         Frontend             Backend
         Next.js                Go
                                 |
                    +------------+------------+
                    |            |            |
                 PubMed      Europe PMC    Crossref
                                 |
                             PostgreSQL

Actual architecture may vary by cloud provider, but responsibilities must remain equivalent.

---

2. ENVIRONMENT

Separate:

development
staging
production

Never use production credentials in development.

---

3. ENVIRONMENT VARIABLES

Document required configuration.

Example categories:

DATABASE_URL
PORT
CORS_ALLOWED_ORIGINS
PUBLIC_APP_URL
PROVIDER_CONFIGURATION
LOG_LEVEL

Actual variable names must follow the implementation.

---

4. SECRET MANAGEMENT

Secrets must come from:

secret manager
deployment platform secrets
secure environment configuration

Never commit:

.env
production credentials
API tokens
private keys
database passwords

".env.example" may contain placeholders only.

---

5. FRONTEND SECRETS

Never expose:

database credentials
backend private secrets
provider private credentials
server tokens

to browser JavaScript.

Only intentionally public configuration may be exposed.

---

6. HTTPS

Production traffic must use HTTPS.

Verify:

HTTP → HTTPS redirect where appropriate
TLS certificate
secure cookies if cookies are used
HSTS where appropriate

Do not claim HTTPS is configured until verified against the deployed environment.

---

7. CORS

Production should allow only the required frontend origin(s).

Avoid:

*

unless there is a documented and justified reason.

---

8. DATABASE NETWORK

PostgreSQL should not be unnecessarily exposed to the public Internet.

Preferred:

Frontend
   ↓
Backend
   ↓
Private DB network

---

9. DATABASE MIGRATION

Migration process:

backup
 ↓
migration
 ↓
verification

For risky migrations:

staging test
 ↓
backup
 ↓
production migration
 ↓
verification

---

10. DATABASE BACKUP

Document:

backup frequency
retention
storage
encryption
restore procedure
owner/responsibility

---

11. HEALTH

Endpoint:

GET /api/v1/health

must be usable for operational checks.

Do not expose sensitive internals in health responses.

---

12. LOGGING

Logs should include useful operational context such as:

timestamp
severity
service
request identifier where available
operation
duration
error category

Avoid sensitive content.

Do not log:

password
API secret
access token
private key
database credentials

---

13. ERROR HANDLING

Production responses should not expose:

stack traces
SQL statements
filesystem paths
internal credentials
provider credentials

---

14. MONITORING

Monitor at least:

uptime
HTTP error rate
latency
CPU
memory
database health
disk/storage

Application-specific signals:

search errors
provider failures
PDF failures
database failures

---

15. ALERTING

Create alerts for meaningful production failures.

Examples:

high 5xx rate
database unavailable
container repeatedly restarting
disk near capacity
persistent provider failures

Avoid excessive alerts that create noise.

---

16. DEPLOYMENT

Recommended:

Build immutable image
 ↓
Tag version
 ↓
Deploy
 ↓
Migration
 ↓
Health check
 ↓
Smoke test

Do not deploy unknown unversioned builds.

---

17. IMAGE SECURITY

Production images should:

use trusted base images
minimize unnecessary packages
run as non-root where practical
avoid embedded secrets

---

18. CONTAINER HEALTH

Verify:

frontend healthy
backend healthy
database healthy

A running process is not necessarily a healthy application.

---

19. RESOURCE LIMITS

Where supported, configure sensible:

CPU
memory
restart policy
connection limits

Avoid unlimited resource consumption.

---

20. EXTERNAL API RESILIENCE

Production behavior remains:

Provider A fails
 ↓
Provider B/C may continue

No provider failure should unnecessarily crash the backend.

---

21. RATE LIMITING

Production search endpoint must remain protected from abusive traffic.

Do not configure limits so aggressively that normal users cannot search.

---

22. PDF

PDF generation must remain:

server-side
controlled
safe

No secrets or internal infrastructure details may appear in generated documents.

---

23. LEGAL FULL TEXT

Production must maintain:

legal access only

No paywall bypass.

No unauthorized article copying.

---

24. DEPLOYMENT VERSIONING

Every deployment should have identifiable:

version
commit
image tag
deployment timestamp

---

25. ROLLBACK

Rollback must be documented for:

application version
configuration
database migration

Never assume database rollback is trivial.

---

26. SMOKE TEST

Production:

homepage
search
results
article
fulltext
original article
PDF

must be verified.

---

27. INCIDENT RESPONSE

Document:

detect
assess
contain
rollback/fix
verify
document

---

28. NO GUARANTEED LATENCY

Third-party services determine part of application latency.

Do not promise:

every search <3 seconds

as an absolute production guarantee.

Instead monitor actual latency.

---

29. PRIVACY

Collect only information required by the product.

Avoid storing unnecessary:

personal information
credentials
sensitive data

---

30. OPERATIONAL DOCUMENTATION

README should contain:

local development
environment setup
migration
testing
deployment
rollback
backup
restore
troubleshooting

---

31. NEXT STATE

After successful deployment:

DEVELOPMENT
    ↓
STAGING
    ↓
PRODUCTION
    ↓
MONITORING
    ↓
MAINTENANCE
