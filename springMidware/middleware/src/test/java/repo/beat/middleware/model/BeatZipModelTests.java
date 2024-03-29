package repo.beat.middleware.model;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

import java.util.Date;

public class BeatZipModelTests {
    @Test
	void testZipModel() {
		String s3Link = "hello:make";
		String zipName = "hello";
		Long zipSize = 188888888883L;
		Date dateAdded = new Date();
		BeatZip beatZip = new BeatZip(zipName, zipSize, s3Link, dateAdded);
		assertEquals(beatZip.getZipName(), zipName);
		assertEquals(beatZip.getZipSize(), zipSize);
		assertEquals(beatZip.getS3Link(), s3Link);
		assertEquals(beatZip.getDateAdded(), dateAdded);

		beatZip.setZipName("world");
		beatZip.setZipSize(188888888883L);
		beatZip.setS3Link("world:make");
		Date tempDate = new Date();
		beatZip.setDateAdded(tempDate);
		assertEquals(beatZip.getZipName(), "world");
		assertEquals(beatZip.getZipSize(), 188888888883L);
		assertEquals(beatZip.getS3Link(), "world:make");
		assertEquals(beatZip.getDateAdded(), tempDate);

    }
}
