import os
import urllib

import petfinder

from google.appengine.api import users

import jinja2
import webapp2


JINJA_ENVIRONMENT = jinja2.Environment(
    loader=jinja2.FileSystemLoader(os.path.dirname(__file__)),
    extensions=['jinja2.ext.autoescape'])

class MainPage(webapp2.RequestHandler):

    def get(self):
        template = JINJA_ENVIRONMENT.get_template('templates/index.html')
        self.response.write(template.render())

class TestPage(webapp2.RequestHandler):

    def get(self):
        # Instantiate the client with your credentials.
        api = petfinder.PetFinderClient(api_key='4b7c674b10cfc0793e935e8abd7346c2', api_secret='15e0e406bc3bcc284773a57957492133')
        pet = api.pet_getrandom()
        self.response.write(pet['name'])


application = webapp2.WSGIApplication([
    ('/', MainPage),
    ('/test', TestPage),
], debug=True)
