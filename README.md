# gf-recaptcha
Idea is to have a script independent of WordPress that uses Captcha 3.x and forks the traffic either to the designated content URL (if human, ie over a certain score we set i the admin back end) For example a python or go (lang) script that sits on a sever (possibly its own subdomain to keep things clean and separated from a load .)

c.example.org (or whatever domain we are running the content on)

So the inbound click URL is:
https://c.example.org/?url=ttps://www.target.org/

If the score is high enough the script redirects to the URL in the query parameter.

If it’s a bot (low score) we redirect to a “we think you are a bot” page. (With no advertising on it) Could also be a page with a ReCaptcha 2.x that then gives them a challenge and then we still let them through if they were a false positive (mis-identified as a bot)



https://example.org/?url=https://www.foo-bar.com would send the human to fooo-bar.


