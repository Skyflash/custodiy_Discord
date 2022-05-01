{{/* Custom Command */}}
{{/* Fornisce le info di base sul progetto */}}
{{/* Utilizzo: -infoita */}}

{{ $icon := (joinStr "" "https://cdn.discordapp.com/icons/" (toString .Guild.ID) "/" .Guild.Icon ".png") }}
{{ $tokens := "1.000.000"}}
{{ $contract := "0x188173379AC8963048Afe01C5d3D5998FEe67254" }}
{{ $buy := "10%"}}
{{ $sell := "10%"}}
{{ $telegram := "https://t.me/custodiy_international"}}
{{ $telegramita := "https://t.me/Custody_app"}}
{{ $telegramchn := "https://t.me/custodiy_cn"}}
{{ $website := "https://www.custodiy.com" }}
{{ $fbpage := "https://www.facebook.com/custodiyofficial"}}
{{ $twitter:= "https://twitter.com/Custodiy1"}}
{{ $instagram := "https://instagram.com/custodiyapp"}}
{{ $discordserver := "https://discord.me/custodiyofficial" }}
{{ $avatar := (joinStr "" "https://cdn.discordapp.com/avatars/" (toString .User.ID) "/" .User.Avatar ".png") }}

{{$embed := cembed 
	"title" "Cosa è CUSTODIY?"
	"description" "**CUSTODIY** è una WEBAPP con quasi tutte le stablecoin in commercio, in ogni valuta.\nApertura smart contract con o senza approvatore.\n\n**IL PROGETTO:**" 
	"color" 0xf07b05 
	"fields" 
		(cslice 
            (sdict "name" "🏦 Monete Totali" "value" $tokens "inline" false)

			(sdict "name" "💰 Tasse di acquisto" "value" $buy "inline" true)
			(sdict "name" "💰 Tasse di vendita" "value" $sell "inline" true)
			
			(sdict "name" "✅ Contratto Verificato" "value" $contract "inline" false)

            (sdict "name" "Sito Web" "value" $website "inline" false)
            (sdict "name" "🇬🇧 Telegram" "value" $telegram "inline" true)
            (sdict "name" "🇮🇹 Telegram" "value" $telegramita "inline" true)
            (sdict "name" "🇨🇳 Telegram" "value" $telegramchn "inline" true)
            (sdict "name" "Twitter" "value" $twitter "inline" false)
			(sdict "name" "Facebook" "value" $fbpage "inline" false)
			(sdict "name" "Instagram" "value" $instagram "inline" false)) 
			"author" 
				(sdict "name" "CUSTODIY - Do It Yourself" "url" "https://www.custodiy.com" "icon_url" "https://images-ext-1.discordapp.net/external/SBstNUrTzLUqZ7Xu4rJurEA_BXy5zl-Bgj3Xxqsbpos/https/cdn.discordapp.com/icons/960427400825540669/67f4ac4e1978fa08a01ca5bffac7a02c.png") 
			"thumbnail" (sdict "url" $icon) 
			"footer" 
				(sdict "text" "Custodiy, il progetto più innovativo della rete!" "icon_url" "https://images-ext-1.discordapp.net/external/SBstNUrTzLUqZ7Xu4rJurEA_BXy5zl-Bgj3Xxqsbpos/https/cdn.discordapp.com/icons/960427400825540669/67f4ac4e1978fa08a01ca5bffac7a02c.png") 
			}}

{{ sendMessage nil $embed }}

{{/* Cancella il trigger che ha scatenato il messaggio */}}
{{ deleteTrigger 0 }}