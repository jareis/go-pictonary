GO=$(shell which go)
RM=$(shell which rm)  -fR
MKDIR=$(shell which mkdir)  -p
CP=$(shell which cp) -Rp
LOCAL_PATH=$(shell realpath .)
TARGET=$(shell realpath ./dist)
TEMPLATES=$(shell realpath ./templates)

build: clean
	$(MKDIR) $(TARGET)
	$(CP) $(LOCAL_PATH)/LICENSE $(TARGET)
	$(CP) $(LOCAL_PATH)/data $(TARGET)
	$(CP) $(LOCAL_PATH)/template $(TARGET)
	$(GO) mod download
	$(GO) mod verify
	$(GO) build -o $(TARGET)

clean:
	$(RM) $(TARGET)
