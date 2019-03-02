# info441finalproject
Joseph Zhang, Kevin Xu, Sunwoo Kang
Info 441 Final Project Proposal
3/1/19
 
Super Smash Bros Tournament Organization Tool
With the recent introduction of Super Smash Bros Ultimate, the competitive scene for the renowned franchise has never been more active. From local to global events, these Smash Bros tournaments are being run daily across the world and they are only gaining more popularity over time. Although these tournaments are being run constantly, it is becoming painstakingly clear that there needs to be more regulation in the organization of these events.
Currently, tournaments use web services like Smash.gg and Challonge to create brackets, which are then overseen and regulated by tournament organizers. Everything is done manually, from the planning to the end of the tournament. This means that tournament organizers have to look through every player’s credentials and place them into a bracket. And then when it comes to the actual tournament, the organizers are running around yelling for people to come play for their turn. Amidst all of this, you factor in human error and inconsistencies, and you get a huge chaotic mess. 
We want to provide a new platform for Super Smash Bros tournaments to smooth out the process for both tournament organizers and players. It will help tournament organizers plan for their events and to help regulate everything on the day of. Players will be able to easily sign up and be alerted of when they need to play, while also making the process of verifying information a lot easier. At the end of the day, everyone involved in the tournament process just want to enjoy playing and watching Super Smash Bros. We want to offer this service to try to minimize the possibility of error and to maximize the amount of time that people get to enjoy the tournament.
Our group honestly just loves the Super Smash Bros series. From the early competitive scene in Super Smash Bros Melee to the recent spike in Super Smash Bros Ultimate, we have been through all of the renditions and have loved this series for a long time. We have all been a part of tournaments both locally and nationally, so we've felt the frustration of having poorly run tournaments firsthand. We believe that this is an area where we can intersect our love of gaming with our abilities as software developers to provide a service that can help optimize something that we, along with many other people, love.
 
 
Technical Description
Architectural Diagram
![architectural diagram](./architecture-diagram.png)
User Cases Table
Priority 
User
Description
P0
As a player
I want to create an account that can save my player information, past tournament standings and seedings.
P1
As a tournament organizer
I want to be able to create a bracket while easily looking at the entrants’ information to take into consideration 
P2
As a tournament organizer
I want to start a new tournament
P3
As a player
I want to be able to sign up for a tournament/look for tournaments to sign up for
P4
As a player
I want to be able to see when and where I play next
P5
As a player
I want to get a notification on when it’s my turn to play
P6
As a player
I want to enter in and confirm the match score after I am done
P7
As a tournament organizer
I want to be able to look at tournament results as the tournament goes on and verify any issues that may come up
 
For each of your user story, describe in 2-3 sentences what your technical implementation strategy is. Explicitly note in bold which technology you are using (if applicable):
Include a list available endpoints your application will provide and what is the purpose it serves. Ex. GET /driver/{id}
Include any database schemas as appendix
User Cases Implementation
P0: Submit post request to /smashgg/users. This adds a record to the users table in the overall tournaments database that contains information about the player. The information can be retrieved with a get request to the same url.
P1: Gather all players entered into tournament from the singular tournament database with a get request to /{tournament}/players and insert into a data structure to track bracket and standing. This insert function will also add the relevant information to the games and pools tables in the singular database table.
P2: Submit a post request to /smashgg/tournaments to create a new tournament. This will create a new singular tournament database dedicated to that tournament and add creator and other tournament organizers to TO table. 
P3: Submit a get request to /smashgg/tournaments to see all available tournaments. Once sign up is verified, a post request is submitted with user id as a query param to {tournament}/players.
P4: Submit a get request to {tournament}/standings to see overall bracket standings, and include a query param to see individual player standing and progress. The get request will retrieve data from the data structure used to hold standings. This request will also retrieve data from the players table joined with the pools table and games table in the singular tournament database.
P5: *Not exactly sure how this is handled. Separate push notification service? Handler that tracks player, game time, and current real time?
P6: Submit a post request to {tournament}/standings with a query param with player id to update player standings. This request will update pools, games, and players tables in the singular tournament database and update the data structure used to hold standings information
P7:  Submit a get request to {tournament}/standings to see overall bracket standings, and include a query param to see individual player standings, pool progress, and game progress. Submit a patch request to the same overall with appropriate query param to verify or solve any issues.
