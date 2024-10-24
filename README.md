# Go Demo Backend 
This is a demo go backend, build with **gin**, and **prisma** with **postgresql** in docker for the database. We will also build the full fledged authentication system from scratch along with the protected routes

### Routes

`/ping` - healthcheck route
`/admin` - protected group route
    `/admin/info` - get the info for currently logged in admin
    `/admin/...` - coming soon
`/auth` - route group for authentication stuff
    `/auth/me` - returns the currently logged in users information
    `/auth/login` - login routes
    `/auth/register` - register routes
    `/auth/logout` - logout routes
`/todo` - todo route group for managing the todo
    `/todo/...` - coming soon
