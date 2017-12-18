# ようこそThe Things Network Japanコミュニティへ!

The Things Networkは、グローバル、オープン、クラウドベースのIoTプラットフォームです。The Things Network日本コミュニティでは、下記の基本原則で運営をおこなっております。

◆**The Things Networkコミュニティの基本原則**
The Things Networkコミュニティの基本原則は次のとおりです。

・The Things Networkマニフェストで言及された原則を支持します
・誰にでもオープンで開かれたコミュニティです
・コミュニティの基盤は、その構成メンバーであり、決してゲートウェイではありません
・定期的な社会的な活動が地域社会の発展の原動力です
・多様性こそが重要です

◆**イニシエータの役割**
コミュニティが出現するためには、コミュニティイニシエータが必要です。The Things Networksは、コミュニティ構築の初期段階でキャンペーンページを作成することを通じてイニシエイタの活動をサポートします。

このページは、特定の分野におけるメンバー、活動、ネットワークの現在の状態に関するすべての関連情報を伝達するためのプラットフォームとして機能します。

地域コミュニティのイニシエータは、常に私たちの完全なサポートを期待することができ、「イニシエータ」というタイトルを付けることができます。その代わりに、イニシアターにマニフェストを尊重し、コミュニティページを最新の状態に保ち、地域ミーティングが組織され、人々が接続できるようにすることを確認します。世界中の人々とこれについて話し合うためのイニシエータスラックチャンネルがあります。

コミュニティが公式に認定されると、イニシエータは「イニシエータ」という称号が永遠に与えられ、誰もあなたからこれを奪うことはありません。

**◆The Things Networkマニフェスト**
マニフェストは、The Things Networkの開始時に作成され、私たちの使命の原則を含んでいます。主な原則は次のとおりです。

「電力を運ぶすべてのモノは、最終的にインターネットに接続されるようになる。」このことを可能にするネットワークを制御することは、世界を制御することになります。我々は、この力が少数の人々、企業、または国に限定されるべきではないと考えている。代わりに、これは誰にも奪われる可能性がない限り、できるだけ多くの人に配布されるべきです。そこで私たちは "The Things Network"を設立しました。The Things Networksは、オープンソース、無料のIoT先駆者として次の性質を有しています。

●The Things Networksは、「IoTノード」と呼ばれるセンサーとアクチュエータを、「IoTゲートウェイ」と呼ばれるトランシーバと、「IoTアクセス」と呼ばれるサーバー群に接続します。

●最初の接続は「無線接続(Over The Air)」、2番目の接続は「ネット接続(Over The Net)」です。これらの概念を分散して実装することを「The Things Network」と呼びます。

●誰でも、自由に "IoT"を設定し、自己または他人が所有する "IoTゲートウェイ"に接続できます。

●誰でも、自由に "IoTゲートウェイ"を設定し、自己または他人が所有する "IoTアクセス"に接続できます。 "IoTゲートウェイ"は、利用可能な最大帯域でのみ制限されるネットワーク中立的な方法で、すべての "IoTノード"へのアクセスを与えます。

●誰でも、自由に「IoTアクセス」を設定し、インターネットからの匿名接続を許可するものとします。"IoTアクセス"は、すべての"IoTゲートウェイ"に利用可能な最大帯域でのみ制限されるネットワーク中立的な方法でアクセスできます。加えて、"IoTアクセス"は、データを配信するために他の"IoTアクセス"との接続を可能とします。

●"無線接続(Over The Air)"と "ネット接続(Over The Net)"用ネットワークは、その通信プロトコルが独自のものでなく、オープンソース、著作権フリーである限り、プロトコル非依存とします。

●「IoTアクセス」または「IoTゲートウェイ」の設置者は、接続されるすべてのデバイスやサーバーに無料でアクセスできます。

●このネットワークを利用する者は、事由を問わず地元の法律によって制限される可能性があること、自己責任でサービスを利用すること、及び「AS-IS」でサービスが提供されること、いつでも解約される可能性があることを、あらかじめ認識するものとします。このネットワークは、不特定利用、限定利用、営利目的、非営利目的、その他の提供方法を問わずオープンに利用できます。「The Things Networks」プロバイダは、ユーザーに制限を設けません。

このマニフェストに署名し、その原則をあなたの能力を最大限に発揮してください。

下記マニフェスト条項を熟読してください[Manifesto](https://github.com/TheThingsNetwork/Manifest).
その他詳細情報に関しましては下記をご参照ください [website](http://thethingsnetwork.org), [forum](http://forum.thethingsnetwork.org/) and [Github](https://github.com/TheThingsNetwork).


# Overview

The Things Network uses the [[LoRaWAN|LoRaWAN/Home]] network technology to provide **low power** wireless connectivity over **long range**.

If we look at The Things Network from a high level, we can distinguish the following components:

[[/uploads/TTN-Overview.jpg]]

* **[[Nodes:|Hardware/Nodes/Home]]** Simple devices that are deployed "in the wild". They can do measurements, collect data or perform actions. Nodes can broadcast or receive small messages (usually about 1/10 the size of an SMS) either periodically (some devices broadcast every couple of minutes, other devices only once in a number of hours).
* **[[Gateways:|Hardware/Gateways/Home]]** Antennas that receive broadcasts from Nodes and send data back to Nodes. Gateways are connected to the Internet and communicate with The Things Network's servers. Gateways have a long range, so they can provide connectivity to nodes that are multiple kilometers away.
* **[[The Things Network Backend|Backend/Home]]** route messages from Nodes to the right Application, and back.
* **Your Application** connects to The Things Network's servers to receive messages from and send messages to your Nodes. What you do with it is entirely up to you!

# Getting Involved

There are many ways to get involved with The Things Network:

* [[Starting or joining a community|Getting_Involved/Community]] doesn't require any technical knowledge. Find an existing local community, or start a new community with other interested people.
* [[Expanding the network|Getting_Involved/Infrastructure]] by contributing infrastructure components, such as gateways (the network needs antennas), a good location for the gateways (maybe you have a nice high building) or servers (to route data between gateways and applications). Any contribution can help your local community and the global TTN community grow the network.
* [[Using the Network|Getting_Involved/UsingTheNetwork]]: **build things** to help solve real world issues with the network. You can help designing use cases, or developing them, or manufacturing them. If possible, share your knowledge and projects with others on any of the communication channels.
* [[Developing the network's backend|Getting_Involved/Development]] is extremely important, especially because the network is growing fast. We need efficient software to route data between gateways and applications. This does require some good programming skills.

**You are the network. Let's build this thing together.**

# Information Sources

Next to this wiki, we have a number of other places where you can find the information you're looking for.

* [The Labs](https://thethingsnetwork.org/labs/)
* [The Forum](https://www.thethingsnetwork.org/forum)
* [Github](https://github.com/TheThingsNetwork) ([backend](https://github.com/TheThingsNetwork/ttn), [examples](https://github.com/TheThingsNetwork/examples), [sdk](https://www.thethingsnetwork.org/docs/applications/sdks.html))
* [Slack](http://slack.thethingsnetwork.org/)
* [Our story](https://medium.com/@wienke/the-things-network-building-a-global-iot-data-network-in-6-months-adc2c0b1ae9b)

>Beware of references to the deprecated REST API (sometimes referred to as "Croft/Jolie", using URLs starting with `thethingsnetwork.org/api`); see [[Connect an Application| Backend/Connect/Application]] to use MQTT instead.

# Social Media

* [Twitter](https://twitter.com/thethingsntwrk)
* [Facebook](https://www.facebook.com/thethingsnetwork)
* [Instagram](https://www.instagram.com/thethingsntwrk)

Local communities often have their own social media pages, you can find these on their community pages.

# External Material

* [[Media Attention|External/Media]]
* [[Blog Posts|External/Blogs]]