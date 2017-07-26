PROJECT_NAME = hownow

.PHONY: test

travis_test:
	cd $(PROJECT_NAME) ; make travis_test
