# **Guide: Configuring Steam Input to disable joystick / controller**

This guide explains how to use set-up and configure Steam Input to ensure the intercepted controllers are not being processed by Train Sim World 5/6/7.

---

## **Step 1: Open Steam and Access Game Properties**

1. Launch **Steam**.
2. Go to your **Library**.
3. Right-click the Train Sim World 5/6/7 and select **Properties**.
  
![Game Properties](https://i.postimg.cc/vHHTxzBp/001-steam-input-guide-open-game.png)  
  
---

## **Step 2: Enable Steam Input**

1. In the **Properties** window, select the **Controller** tab.
2. Under **Override for [Game Name]**, select **Enable Steam Input**.
3. Close the properties window.
  
![Force Enable Steam Input](https://i.postimg.cc/xTCj11pS/002-force-enable-steam-input.png)  
  
> **Note:** This ensures Steam Input is always active for Train Sim World.

---

## **Step 3: Calibrate an Unrecognized Controller (if needed)**

If your controller is unrecognized by Steam (which is often the case for custom joysticks):

1. Go to **Steam → Settings → Controller**.
2. Connect your controller.
3. Select "Configure Controller" or "Calibration" for the controller you are trying to intercept.
4. In the calibration window, you will only need to configure the required "A" and "B" buttons; everything else can be skipped by re-hitting "A" repeatedly.

---

## **Step 4: Apply the Community Layout**

1. With your controller connected, click the **Controller Configuration** button in your game’s Library page.
2. Instead of browsing, use the direct link for the "Disabled Gamepad" layout:

**Train Sim World 5**  
[steam://controllerconfig/2967990/3576092503](steam://controllerconfig/2967990/3576092503)  
  
**Train Sim World 6**  
[steam://controllerconfig/3656800/3576139582](steam://controllerconfig/3656800/3576139582)  
  
![Apply Gamepad Layout](https://i.postimg.cc/d3X0YMYC/003-apply-layout.jpg)  
  
3. Steam will open the layout and ask to **Apply Configuration**. Click **Apply**.

> **Note:** This particular template is designed to **disable all controller input for the game**, allowing another software to handle the controller.

---

## **Step 5: Verify the Layout**

1. Launch the game.
2. Test that the game ignores controller input
3. If the controller is still active in-game, double-check **Controller Configuration** and ensure the layout is applied.
