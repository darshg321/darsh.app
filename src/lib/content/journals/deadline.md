---
title: "Deadline"
github: "https://github.com/darshg321/Deadline"
description: "putting the dead in deadline"
created_at: "2026-08-19"
---

# Aug 19th: research and plan
![image](https://cdn.hackclub.com/01a01cb1-bc7f-7e62-b3ff-809649aa2bf0/paste-1787187805377.png)
made the plan! its an overengineered mechanical progress bar: putting the dead in deadline
- ball screw driven by a stepper that gets closer to a deadline set by an mcu
by percentage
- once it gets to the end, it hits a lever that drops a little 3d printed guillotine
on a stickman

made a diagram because i think thats the best way to plan a project especially a cad. had to research a bunch of 
components for this, so i ended up also basically making a BOM while doing the thing. pasted in references to all
the images, to draw this diagram and then labelled everything. 

what i learned was a few things like that lead screws and ball screws are different things, where the ball screws 
are actually better in like every way but theyre more expensive. i also learned the drv stepper driver is way better
than other stepper drivers, and that there are two versions of the nema17 but i already got an aliexpress bundle thats
the cheaper and worse nema17 motor. i learned how coupling works, thats how the ball screw is driven. 

tasks to do:
- find imports for the ball screw and as many components as i can
- create 3d prints for both the load and the guillotine and person, and mounting for the motor
- create an assembly that moves
- render this assembly
- create a wiring diagram
- create the code to pull the deadline with wifi (can be my google calendar), and move the stepper in line with it

time spent: 1h37m (https://lapse.hackclub.com/timelapse/olxA_3zSMcl1)

# Aug 20th: cad!
started the cad! i first found a bunch of parts to import. firstly i had to figure out how the coupler dimensions work, 
which is two different sizes based on the shafts. 

![image](https://cdn.hackclub.com/01a01e0d-39db-7722-81b7-d46058c458ab/paste-1787210578744.png)

![image](https://cdn.hackclub.com/01a01e0e-17bf-7ff3-968b-6f75006e9120/paste-1787210635724.png)

imported some stuff in

![image](https://cdn.hackclub.com/01a01e0d-8d0a-7d48-98c9-f1e1a5efbb6a/paste-1787210600121.png)

figured out how screw relation works in onshape

![image](https://cdn.hackclub.com/01a01e0e-d2d1-795a-bfb2-e3ef34cb1a44/paste-1787210683561.png)

![image](https://cdn.hackclub.com/01a01e0e-8bc6-7f84-9f54-fa8c37946c8b/paste-1787210665422.png)

![image](https://cdn.hackclub.com/01a01e0f-0af7-7dd9-845c-7de3ce35ae3d/paste-1787210697982.png)

worked on the assembly further

![image](https://cdn.hackclub.com/01a01e0f-48ff-74e1-af1a-7bb44b0b1dc9/paste-1787210713846.png)

made a little stickman

![image](https://cdn.hackclub.com/01a01e0f-b58a-7649-a714-35a751692942/paste-1787210741640.png)

time spent: 2h3m (https://lapse.hackclub.com/timelapse/BwpDlXNMCpqm)

# Aug 20th: finish project!

i finished the cad!

i realized i needed a way to prevent the load (guillotine) from rotating, so i added slider rods to the side of the holder, to keep it in place while the ball screw rotates

![image](https://cdn.hackclub.com/01a01fbc-96a0-75a1-b483-3e1bece20142/paste-1787238847971.png)

then i added the stickman

![image](https://cdn.hackclub.com/01a01fbc-f51f-7aa1-9d53-6007918c50b2/paste-1787238871527.png)

i had to fiddle around with assembly mates a lot because i couldnt get the sliding part to not rotate while the ball screw was rotating, (it would create overdefined)
so i had to settle with just fastening it and fastening the ball screw holder separately

![image](https://cdn.hackclub.com/01a01fbc-3fc1-718c-838c-5c7e2484e8bf/paste-1787238825507.png)

then i created the guillotine
![image](https://cdn.hackclub.com/01a01fbd-6c01-7db8-aece-1b68a1929628/paste-1787238902481.png)

i created a wiring diagram for the esp32s3, nema17, drv8825, and as5600. it seems a lot of pins dont actually need to be used
![image](https://cdn.hackclub.com/01a01fbd-dd9a-70a9-997d-4a2d5792ac9c/paste-1787238931680.png)

created a BOM and had to remove the screws i wanted because slightly too expensive
![image](https://cdn.hackclub.com/01a01fbd-9d42-7575-8e5e-194e174d8341/paste-1787238915263.png)

finally im writing this journal
![image](https://cdn.hackclub.com/01a01fbe-1359-7689-ae9d-721e08891ac9/paste-1787238945531.png)

time spent: 1h47m (https://lapse.hackclub.com/timelapse/IHWwiemtrXS8)

