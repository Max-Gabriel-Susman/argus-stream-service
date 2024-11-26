
# deprecated
compile: 
	g++ main.cpp

# deprecated
run: 
	./a.out

test-gstreamer:
	gst-launch-1.0 videotestsrc ! videoconvert ! autovideosink