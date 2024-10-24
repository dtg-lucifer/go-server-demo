# Go Demo Backend 
This is a demo go backend, build with **gin**, and **prisma** with **postgresql** in docker for the database. We will also build the full fledged authentication system from scratch along with the protected routes

### Routes

`/ping` - healthcheck route <br>
`/admin` - protected group route <br>
- `/admin/info` - get the info for currently logged in admin <br>
- `/admin/...` - coming soon <br>
`/auth` - route group for authentication stuff <br>
- `/auth/me` - returns the currently logged in users information <br>
- `/auth/login` - login routes <br>
- `/auth/register` - register routes <br>
- `/auth/logout` - logout routes <br>
`/todo` - todo route group for managing the todo <br>
- `/todo/...` - coming soon <br>
