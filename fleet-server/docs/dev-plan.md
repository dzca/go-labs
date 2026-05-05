To build a multi-tenant fleet system, you’ll use Keycloak as the central Identity and Access Management (IAM) engine and a Go backend to manage your business data (fleets, devices, drivers). [1]

## 1. Keycloak Architecture for Multi-Tenancy

You have three main strategies for multi-tenancy in Keycloak:

- Realm-per-Tenant: Highest isolation. Each fleet gets its own realm. Best for "white-label" systems where each fleet needs its own branding or login methods.
- Group/Attribute-based (Single Realm): Users share a realm but are tagged with a fleet_id attribute. This is easier to manage at scale (thousands of fleets).
- Organizations (New in v25+): A native "Organization" concept that groups users and identity providers within a single realm. [2, 3, 4, 5, 6, 7]

---

## 2. Required API Implementation

You will need approximately 10–12 core API endpoints to handle onboarding and management.

## A. Onboarding & Fleet Management (Super Admin only)

1.  POST /api/v1/fleets: Onboards a new fleet. This triggers the creation of a tenant group in Keycloak or a new Realm.
2.  POST /api/v1/fleets/{id}/activate: Activates/Deactivates the fleet status.
3.  GET /api/v1/fleets: Lists all fleets (with filtering for Super Admins). [8, 9]

## B. User & Manager Management (Fleet Admin/Manager)

1.  POST /api/v1/users/managers: Adds a fleet manager and creates their account in Keycloak using the [Keycloak Admin REST API](https://www.keycloak.org/docs-api/latest/rest-api/index.html).
2.  POST /api/v1/users/drivers: Adds a driver. Your Go backend should call Keycloak to create the user, then save the Keycloak user_id and fleet_id in your PostgreSQL.
3.  DELETE /api/v1/users/{id}: Removes a user from both Keycloak and your DB. [10, 11]

## C. Device Management (Fleet Admin/Manager)

1.  POST /api/v1/devices: Registers a new device (e.g., GPS tracker) and assigns it to a fleet.
2.  PATCH /api/v1/devices/{id}/link: Links a specific driver to a device for a shift.
3.  GET /api/v1/devices: Lists devices filtered by the logged-in user's fleet_id.

## D. Utility/Auth

1.  GET /api/v1/me: Returns the current user's profile, including fleet membership extracted from the Keycloak JWT. [12]

---

## 3. Implementation Steps

1.  Keycloak Setup:

- Create a "Confidential Client" for your Go backend with service-accounts enabled so it can manage users via API.

2.  Go Backend Onboarding:

- Step 1 (Fleet Creation): Create a record in your fleets table.
  - Step 2 (IAM Link): Use a library like [Gocloak](https://github.com/Nerzal/gocloak) to create a group or organization in Keycloak named after the fleet.

3.  User Provisioning:

- When adding a driver, your Go code calls Keycloak's /admin/realms/{realm}/users to create the account.
  - Assign the driver to the fleet's Keycloak group and set a fleet_id custom user attribute.

4.  Token Validation (Middleware):

- In every Go API call, extract the JWT from the Authorization header.
  - Verify the signature using Keycloak’s JWKS endpoint.
  - Extract the fleet_id from the token claims to ensure the user can only see their own fleet's devices or drivers. [10, 13, 14, 15]

## 4. Components Diagram

- Keycloak (Docker): Handles login, user sessions, and JWT issuance.
- PostgreSQL: Stores your fleets, devices, and drivers (link drivers to Keycloak via keycloak_uuid).
- Go Backend (Gin/Echo): Contains your business logic and Keycloak Admin client.
- Flutter App: Authenticates against Keycloak, sends the JWT to the Go Backend. [12, 16]

Would you like a sample JSON payload for the fleet manager creation or a Go code snippet using Gocloak to link a driver?

[1] [https://www.youtube.com](https://www.youtube.com/watch?v=DNq51wWw3F4&t=151)
[2] [https://medium.com](https://medium.com/swlh/using-keycloak-for-multi-tenancy-with-one-realm-7be81583ed7b)
[3] [https://www.reddit.com](https://www.reddit.com/r/KeyCloak/comments/1apubm9/keycloak_for_a_multitenant_app_how_to_design_it/)
[4] [https://www.youtube.com](https://www.youtube.com/watch?v=ZTFlc-3pG1M&t=10)
[5] [https://medium.com](https://medium.com/swlh/using-keycloak-for-multi-tenancy-with-one-realm-7be81583ed7b)
[6] [https://www.youtube.com](https://www.youtube.com/watch?v=tY06l4KRHKk)
[7] [https://www.youtube.com](https://www.youtube.com/watch?v=tY06l4KRHKk)
[8] [https://aws.amazon.com](https://aws.amazon.com/blogs/compute/managing-multi-tenant-apis-using-amazon-api-gateway/)
[9] [https://medium.com](https://medium.com/@nivas.ganesan/multitenantcy-keycloak-loa-with-spring-boot-3-2-1-9c2b8bee8d98)
[10] [https://www.reddit.com](https://www.reddit.com/r/golang/comments/1efwao2/has_anyone_integrated_keycloak_in_a_golang_app/)
[11] [https://www.youtube.com](https://www.youtube.com/watch?v=eZYGLuUrUp4)
[12] [https://dev.to](https://dev.to/zrouga/build-a-secure-multi-tenant-sso-system-with-keycloak-go-react-step-by-step-guide-218m)
[13] [https://medium.com](https://medium.com/@rizkysr90/getting-started-with-oauth-2-0-in-golang-using-keycloak-8e61fdf3620b)
[14] [https://medium.com](https://medium.com/@vgzxkgmrpn/securing-rest-api-endpoints-with-oauth-2-0-and-keycloak-client-scopes-b5979702472a)
[15] [https://github.com](https://github.com/ksingh7/keycloak-go-demo)
[16] [https://www.keycloak.org](https://www.keycloak.org/docs/25.0.6/securing_apps/index.html)
