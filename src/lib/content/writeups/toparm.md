

## What is TopArm?

TopArm is a fuse of a top hat and a robotic arm. It started as a random idea I had after seeing some 3d-printed robot
arms with micro servos, getting me into robotic joints and how they work, and additionally finally using my 3d printer
for a project.

---

## Supercon 2025


I showed off TopArm at Supercon 2025, where it immediately got the attention of most people walking by, showing me how
even the most simply things can be impressive to people. This project wasn't exactly technical, it was only some
3d-printed parts and code to run a few motors, strapped to a hat, but somehow people really did like it.

It ended up:

**Winning 3rd place in the costume contest**, winning me an unreasonably expensive lab-grade waveform generator

**Impressing a lot of serious hardware people**, including the Framework CEO

**Not breaking**, despite a fall and completely untightened nuts and bolts

The highlight of the conference was definitely walking around, talking to all of these people about their own projects
and what they do for work, and of course leading them back to show them my project.

---

## How It Works

### Control

TopArm is controlled with a **joystick**, making it usable while walking around a conference floor, grabbing props,
or just gesturing dramatically while talking.

### Electronics

* **Microcontroller:** Arduino
* **Actuation:** SG90 micro servos
* **Servo Control:** PCA9685 servo driver

### Mechanics

The arm is relatively lightweight, and with the help of thumbtacks on the base, it was manageable to wear while walking
around. All other components are held together by zip tie, and some tape that definitely did not work effectively since
I ended up having to retape every component multiple times throughout each day. There were multiple connection parts
that I didn't have, and so ended up 3d printing which worked surprisingly well. What didn't work well were the screws,
I'm not completely sure if I just needed to tweak tolerances of the 3d print or get different sized screws, but they
were loosening, and while it didn't fall apart ever, I had to constantly tighten each screw by hand.

---

## Lessons Learned

* Simplistic, working functionality is usually still good to show off
* Joins and motors aren't as scary to understand as I once thought

While I won't be working on this project any longer, I've gotten much closer to designing robots, and my new interest
for the past few weeks has been actuators and quadruped limbs, and so this project, along with some very interesting
videos, and a great conference talk
[(What Makes a Robot Feel Alive? by Cyril Engmann)](https://docs.google.com/presentation/d/1S253Vh8X4zbfiBJ7xzXYNUvq2RK8Lu6_zbJwN3hTrp0/edit?slide=id.p#slide=id.p)
has inspired me to continue forward with my next project, Psidog.