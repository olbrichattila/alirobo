## Alibaba cloud game challenge

**This is a submission for the [Alibaba Cloud](https://int.alibabacloud.com/m/1000402443/) Challenge: [Build a Web Game](https://dev.to/challenges/alibaba).**

---

## First, The credits:
For images credit goes to https://opengameart.org/, the open source game art community, I could not do it in time without them!

----

## What I Built
In this game, you play as a shape-shifting ethical hacker robot on a mission assigned by your boss. Your objective is to use social engineering tactics to infiltrate eight different server rooms located in a secret underground facility.

Navigate the maze of tunnels and ladders using the arrow keys. Your first goal is to locate all the office rooms.
If an office has an employee moving around, you must first shape-shift into a human form by pressing the S key. Be aware that if you move afterward, you will revert to your original robot form. Wait for the employee to come close, then press the Shift key to interact. This represents a social engineering attempt where you claim to be an engineer who forgot their access badge. If successful, you’ll receive a new badge, and a corresponding server room icon will appear, indicating the type of server room the badge grants access to.

Some offices won’t have employees but will contain a computer on a desk. Approach the computer and press the Space bar to collect a password. This simulates the ethical hacker discovering an unattended, unlocked machine and retrieving credentials from it.
Once you’ve collected all the badges, they will begin to flash, signaling that you are now ready to locate and access the server rooms on levels -5 and -6.

Each server room requires the correct badge to enter. The type of server room is indicated by an icon displayed near the door, representing different Alibaba services. You must recognize the correct icon and use the corresponding badge. You have three attempts per room—after that, you’ll lose a life.
Successfully entering a server room allows you to continue to the next one.

After infiltrating all server rooms, return to the starting room, where your boss is waiting. If all server rooms have been accessed, your mission is complete. This final step symbolizes reporting your findings to your boss so they can initiate appropriate security training.


If for the specific room available and you're inside a room and unsure what to do, press the H key to receive a hint—provided the room has special rules or tasks.

---

## Demo
(I've registered a domain for -> pointing to Alibaba ECS)
Please note that the game is designed to work on a computer; mobile devices are not supported, and some old machines may not work due to a lack of Opengl support.

> **Please give a few minutes to load!!!**  
[https://www.alirobo.fun/](https://www.alirobo.fun/)


---

## Screenshots

![Opening page](https://raw.githubusercontent.com/olbrichattila/alirobo/main/screenshots/sc1.png)
![Opening page](https://raw.githubusercontent.com/olbrichattila/alirobo/main/screenshots/sc2.png)
![Opening page](https://raw.githubusercontent.com/olbrichattila/alirobo/main/screenshots/sc3.png)
![Opening page](https://raw.githubusercontent.com/olbrichattila/alirobo/main/screenshots/sc4.png)
![Opening page](https://raw.githubusercontent.com/olbrichattila/alirobo/main/screenshots/sc5.png)
![Opening page](https://raw.githubusercontent.com/olbrichattila/alirobo/main/screenshots/sc6.png)
![Opening page](https://raw.githubusercontent.com/olbrichattila/alirobo/main/screenshots/sc17.png)
![Opening page](https://raw.githubusercontent.com/olbrichattila/alirobo/main/screenshots/sc18.png)
![Opening page](https://raw.githubusercontent.com/olbrichattila/alirobo/main/screenshots/sc19.png)
![Opening page](https://raw.githubusercontent.com/olbrichattila/alirobo/main/screenshots/sc20.png)

---

## Alibaba Cloud Services Implementation

My game is written entirely in Golang and compiled to WebAssembly (WASM) to run directly in the browser. Since I work with Golang professionally on backend systems, I thought—why not build a game with it too? To support the gameplay, I wrote a lightweight Golang-based API that handles storing and retrieving scores, including listing the top 10 leaderboard entries.


**Planned Architecture Using Alibaba Cloud Services**
Initially, the infrastructure I planned for the game made use of the following Alibaba Cloud services:

- **Object Storage Service (OSS)**
Used to host static assets such as the HTML, WASM, JavaScript, and image files.
Why OSS? OSS offers a scalable, high-performance storage option and can easily integrate with a CDN for faster delivery worldwide.
Integration Experience: The setup wasn’t as seamless as expected. OSS forces static files like HTML to be downloaded instead of rendered in the browser unless a custom domain is configured with HTTPS, which wasn’t obvious during setup. I had to dig through documentation to understand this behavior.
Challenge: I didn’t want to register a custom domain just for HTTPS, so this created friction during deployment.

- **CDN + OSS**
I had planned to serve images via the Alibaba Cloud CDN connected to OSS.
Benefit: Faster asset delivery with minimal latency.
Challenge: Due to the OSS hosting limitation mentioned above, I didn’t proceed with the CDN as originally planned.

- **Function Compute (FC)**
The Golang API was designed to run as a Serverless Function Compute.
Why Function Compute? It was a good fit for the small, stateless HTTP endpoints of the score API. Serverless meant simplified scaling and reduced maintenance overhead.
Challenge: I encountered limitations with the free tier or trial—either I didn’t receive the trial credit or it was extremely restricted. Cost estimates showed that using FC + managed services would exceed hundreds of dollars/month, which is not feasible for a personal, non-commercial game. This issue was echoed by several developers in Alibaba Cloud’s discussion forums.

- **ApsaraDB for PostgreSQL**
Intended for storing game scores in a managed relational database.
Why ApsaraDB? It offered managed backups, security, and scalability, aligning with my professional stack.
Challenge: Same as above—cost and trial limitations made it impractical to use for a personal project.

- **Elastic Compute Service (ECS)**
As a backup plan, I experimented with ECS to run everything in a containerized setup.
Challenge: Strangely, the ECS instance I tested performed significantly slower compared to equivalent instances on another cloud provider—around 20x slower. I couldn’t identify the cause, but it made testing very sluggish and affected the user experience.

---

- **Final Setup Due to Cost Constraints**
Due to the challenges around cost and configuration, I settled on a minimal paid setup using:

**Alibaba Cloud ECS*** (Elastic Compute Service) – $11/month for the smallest instance

Hosting the entire stack:
- WebAssembly game (HTML, JS, WASM, images)
- Golang HTTP API
- PostgreSQL database (all in a single Docker container)

This allowed me to maintain full control over hosting while staying within budget. It also made HTTPS straightforward to configure with a reverse proxy.

---

## Game Development Highlights
Interesting Aspects of Development
I decided to move forward with this game at the last minute after the deadline was extended, so the time pressure is definitely reflected in the code — which is publicly available here: [https://github.com/olbrichattila/alirobo](https://github.com/olbrichattila/alirobo).

This was my first game written in Go, and it was a great opportunity to experiment with the Ebiten framework. Throughout development, I learned a lot about rendering, sprite handling, game loop logic, and behavioral interactions within a game context.

One unexpected challenge I ran into was compatibility issues with older Windows machines, particularly due to OpenGL limitations that Ebiten relies on. It was an eye-opener in terms of how graphics APIs can affect cross-platform support.

Despite the time crunch and hurdles, I'm proud of how much I was able to build in such a short time and how I pushed Go beyond backend development into the realm of browser-based gaming.
