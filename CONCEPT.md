# Go Hot Potato Concept and Mechanics

The main idea behind this game is to show how concurrency works in `Golang` using a game. The games essential premise is to have the *Potato* passed around by the players until the timer is reached and the last person to hold it would be exiting the game as a looser.

### Rules of the Potato Game

- [ ]  For the game to start there should be five *Players* up until then their would be a lobby where they would stay.
- [ ]  A *Player* is said to have lost if they have the *Potato* when the timer is done.
- [ ]  The game is played until one Player remains so five rounds in total.

### Theme

“Hot Potato” style game where players pass a digital potato under a timer.

### Players Interactions and Actions

- Log to play
- Pass the Potato
- Ready/Unready
- Leave Lobby
- Challenge Pass
- Passing to Self is Invalid

### Event Flow

- Players get into a system get assigned a Go-routine.
- If no Players are available they get to wait in the lobby.
- Once Players required amount i.e. Five is reached game starts.
- The Potato gets randomly assigned to Players and they should pass it out until the time is out.
- The last Player to have the Potato when the time ticker is done is eliminated.

Summary: **Lobby → Ready Check → In-Game → Round End → Elimination → Next Round**.

### Places for Concurrency Usage

- Current time is passed for each Go routine
- The Connections of each Go-routine is kept alive
- The Potato is passed through channels to each Go-routine at random.
- Each Go-routine would be able to pass the Potato to another Go-routine unless the limited time is up.

### Core Structs

<aside>
🎮

Player{

ID        string
Conn      net.Conn
IsReady   bool
IsActive  bool
LastPass  time.Time

}

</aside>

<aside>
🥔

Potato{

HolderID  string
PassCount int
Timer     *time.Timer

}

</aside>

<aside>
🕹️

Game-State{

Round       int
Players     map[string]*Player
PotatoChan  chan string   // holds next holder’s ID
StateMutex  sync.Mutex

}

</aside>
