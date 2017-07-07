# Troubleshooting

## I have a problem

- Check the TTN status page (under construction), the Forum, the `#ops` channel on Slack or our Twitter account to see if there are any service disruptions.

## I am no longer receiving data in my Application

- Check the live traffic for your device in the "Data" tab or using `ttnctl subscribe`.
	- If you don't see anything showing up, there might be a problem with your access keys, or data is not received by the network.
- Check the status of your device on the Console or using `ttnctl devices info [DevID]`.
	- If the **Status** or **Last Seen** shows a time that is long ago, data is indeed not received by the network

## The network is no longer receiving data from my Node

- Check that your device uses OTAA, and not ABP.
	- If your device uses ABP, you might want to change this, or check the Frame Counters.
- Check that you have coverage in the area of your Node
	- If you do not own a gateway in the area, the Node may not have coverage or someone elses Gateway may be unavailable.

## I am no longer seeing traffic for my Gateway

- Check the status of your gateway on the Console and using `ttnctl gateways status [GatewayID]`
  - If the **Last Seen** on the Console shows a time that is long ago, your gateway may be offline, or there is a problem with the NOC
	- If the **Last Seen** in `ttnctl` shows a time that is long ago, your gateway is likely offline
- Check the logs of your gateway to verify that it is still working
- Check internet traffic of your gateway (using `tcpdump`) to verify that it is transmitting data

## I still have a problem

- Notify the `#ops` channel on Slack. When multiple people have the same problem, we can investigate deeper.
	- `Possible problem in {eu,us,asia,...}: No longer receiving data for {application+device,gateway,...} {id} since {time}`
	- `Possible problem in {eu,us,...}: {MQTT/API/Console} seems to be down`

## Support Ticket

- If you seem to be the only one experiencing this problem, you can submit a support ticket with one of the (commercial) support providers. 
- Please provide the following information:
	- What do you want to do?
	- What steps did you already take?
	- What went wrong?
	- Please provide as many IDs as possible: Gateway IDs, Device IDs, Application IDs, Usernames
		- Please send these IDs as text, not as screenshots
	- What kind of node/gateway/application? What kind of software is running on it? What version?
	- When did it stop working?
	- What do your log files look like?
  - Do you have screenshots?