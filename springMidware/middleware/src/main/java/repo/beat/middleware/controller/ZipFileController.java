package repo.beat.middleware.controller;

import java.io.File;
import java.util.Date;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import repo.beat.middleware.model.BeatZip;
import repo.beat.middleware.service.ZipFileDBServiceImpl;
import repo.beat.middleware.service.ZipFileS3ServiceImpl;

import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;


@RestController
@RequestMapping("/api/zip")
public class ZipFileController {
    @Autowired
    private ZipFileDBServiceImpl zipFileDBService;

    @Autowired
    private ZipFileS3ServiceImpl zipFileS3Service;

    @PostMapping("/upload")
    public void uploadHandler(@RequestParam String bucketName, @RequestParam String objectKey, @RequestParam String filePath) {
        String s3Link = bucketName + ":" + objectKey;
        zipFileS3Service.uploadZipFile(bucketName, objectKey, filePath);
        BeatZip beatZip = new BeatZip(objectKey, new File(filePath).length(), s3Link, new Date());
        zipFileDBService.saveZipFile(beatZip);
    }

    @GetMapping("/download")
    public void downloadHandler(@RequestParam String bucketName, @RequestParam String objectKey, @RequestParam String filePath) {
        zipFileS3Service.downloadZipFile(bucketName, objectKey, filePath);
    }

    @DeleteMapping("/delete")
    public void deleteHandler(@RequestParam Long id) {
        BeatZip beatZip = zipFileDBService.getZipFileById(id);
        String[] s3Link = beatZip.getS3Link().split(":");
        String bucketName = s3Link[0];
        String keyName = s3Link[1];

        zipFileS3Service.deleteZipFile(bucketName, keyName);
        zipFileDBService.deleteZipFile(id);
    }

}

