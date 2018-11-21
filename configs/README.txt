setup:

	1. download [prebuild Binary] .
		git clone  https://github.com/annieweng/oauth2_proxy.git
		download corresponding precompiled bin/*.tar.gz for your OS.  

	2. register an Oauth Application with DSRA Oauth2 provider at https://oauth2-server.dsra.io
		take a note of client ID/secret generated from registration. make sure your redirect URL match exactly to where you running the oauth2_proxy from. 
		typically this will be http[s]://hostname/oauth2/callback

	3. configure Nginx to run https and proxy all traffic coming to https://hostname to oauth2_proxy:

	⁃	to generate a self-sign cert for test(this is only needed if you don’t already have SSL configured):
		mkdir -p /etc/nginx/ssl
		openssl req -x509 -newkey rsa:2048 -keyout key.pem -out cert.pem -days 360 -nodes
		this will generate key.pem and cert.em files.

	⁃	add proxy config to oauth2_proxy running at http://localhost:4180(default port)

	⁃	 here is what the /etc/nginx/nginx.conf will look like:

		server{
				listen  443 ssl;
                server_name    localhost.com;
				ssl     on;
                ssl_certificate         /etc/nginx/ssl/cert.pem;
                ssl_certificate_key     /etc/nginx/ssl/key.pem;
                ssl_protocols   TLSv1 TLSv1.1 TLSv1.2;

		location / {
    			proxy_pass http://127.0.0.1:4180;
   		 	proxy_set_header Host $host;
    			proxy_set_header X-Real-IP $remote_addr;
    			proxy_set_header X-Scheme $scheme;
    			proxy_connect_timeout 1;
    			proxy_send_timeout 30;
    			proxy_read_timeout 30;
 		 }
		}	
	    it's also available at configs/ directory.

	4. configure oauth_proxy to run at http://localhost:4180
	⁃	    copy configs/oauth2_proxy.cfg from git source tree to /etc/oauth2_proxy directory
	   	  change clientid, secret, cookie secret value as need

	⁃	  copy configs/oauth2_proxy.conf to /etc/supervisor/conf.d/
	           supervisorctl reread; supervisorctl update;


		
	
	5. navigate to https://$hostname, it will automatically redirect you to login page and take you to oauth2 provider to start authentication process. 
	oauth2_proxy will redirect you back to your upstream client(configured in oauth2_proxy.cfg) once the login process is successfully completed.
		



